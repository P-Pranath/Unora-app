import {MMKV} from 'react-native-mmkv';
import {StateStorage} from 'zustand/middleware';

// Create MMKV instance
export const storage = new MMKV();

/**
 * Zustand-compatible MMKV storage adapter
 * Use this with Zustand's persist middleware
 * Zustand serializes to JSON strings, so we handle strings only
 */
export const mmkvStorage: StateStorage = {
  getItem: (name: string): string | null => {
    const value = storage.getString(name);
    return value ?? null;
  },
  setItem: (name: string, value: string): void => {
    storage.set(name, value);
  },
  removeItem: (name: string): void => {
    storage.delete(name);
  },
};

/**
 * Helper functions for direct MMKV usage
 */
export const mmkvHelpers = {
  /**
   * Set a string value
   */
  setString: (key: string, value: string): void => {
    storage.set(key, value);
  },

  /**
   * Get a string value
   */
  getString: (key: string): string | undefined => {
    return storage.getString(key);
  },

  /**
   * Set a number value
   */
  setNumber: (key: string, value: number): void => {
    storage.set(key, value);
  },

  /**
   * Get a number value
   */
  getNumber: (key: string): number | undefined => {
    return storage.getNumber(key);
  },

  /**
   * Set a boolean value
   */
  setBoolean: (key: string, value: boolean): void => {
    storage.set(key, value);
  },

  /**
   * Get a boolean value
   */
  getBoolean: (key: string): boolean | undefined => {
    return storage.getBoolean(key);
  },

  /**
   * Set a JSON object
   */
  setObject: <T>(key: string, value: T): void => {
    storage.set(key, JSON.stringify(value));
  },

  /**
   * Get a JSON object
   */
  getObject: <T>(key: string): T | undefined => {
    const value = storage.getString(key);
    if (value) {
      try {
        return JSON.parse(value) as T;
      } catch {
        return undefined;
      }
    }
    return undefined;
  },

  /**
   * Delete a key
   */
  delete: (key: string): void => {
    storage.delete(key);
  },

  /**
   * Check if key exists
   */
  contains: (key: string): boolean => {
    return storage.contains(key);
  },

  /**
   * Get all keys
   */
  getAllKeys: (): string[] => {
    return storage.getAllKeys();
  },

  /**
   * Clear all storage
   */
  clearAll: (): void => {
    storage.clearAll();
  },
};

export default storage;
